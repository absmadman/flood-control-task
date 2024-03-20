package floodcontrol

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type FC struct {
	Ctx    context.Context
	Client *redis.Client
}

// TokenBucket структура ведра с токенами
type TokenBucket struct {
	tokens         int
	maxTokens      int
	refillTokens   int
	lastRefillTime time.Time
	timeToRefill   int
}

// RedisVal структура для удобства хранения полей в Redis
type RedisVal struct {
	Tokens         int       `json:"tokens"`
	LastRefillTime time.Time `json:"lastRefillTime"`
}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}

// User структура для хранения клиентов в базе которая содержит время последнего пополнения токенов
// и их количество
type User struct {
	UserID         int64
	tokens         int
	lastRefillTime time.Time
}
