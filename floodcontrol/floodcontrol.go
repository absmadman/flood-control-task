package floodcontrol

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

// NewFC конструктор для FloodControl структуры
func NewFC(ctx context.Context) *FC {
	rDb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return &FC{
		Ctx:    ctx,
		Client: rDb,
	}
}

// NewUser конструктор для User'a
func NewUser(UserID int64, tokens int, lastRefillTime time.Time) *User {
	return &User{
		UserID:         UserID,
		tokens:         tokens,
		lastRefillTime: lastRefillTime,
	}
}

// GetData получает данные из Redis и складывает их в структуру User с помощью json.Unmarshall
func (fc *FC) getData(userID int64) (*User, error) {
	var rv RedisVal
	buf, err := fc.Client.Get(context.Background(), string(userID)).Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &rv)
	if err != nil {
		return nil, err
	}
	return NewUser(userID, rv.Tokens, rv.LastRefillTime), nil
}

// PutData формирует []byte из User'a с помощью json.Marshal и складывает данные в Redis
func (fc *FC) putData(user *User) error {
	buf, err := json.Marshal(&RedisVal{
		Tokens:         user.tokens,
		LastRefillTime: user.lastRefillTime,
	})
	if err != nil {
		return err
	}
	err = fc.Client.Set(context.Background(), string(user.UserID), buf, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// Check использует Token Bucket для проверки пользователей на флуд
func (fc *FC) Check(ctx context.Context, userID int64) (bool, error) {
	user, err := fc.getData(userID)
	if err != nil {
		if err = fc.putData(NewUser(userID, 5, time.Now())); err != nil {
			return false, err
		}
	} else {
		tb := NewTokenBucket(user.tokens, user.lastRefillTime)
		if tb.request() == false {
			return false, nil
		}
		if err = fc.putData(NewUser(userID, tb.tokens, tb.lastRefillTime)); err != nil {
			return false, err
		}
	}
	return true, nil
}
