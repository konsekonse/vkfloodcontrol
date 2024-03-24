package main

import (
	"context"
	"sync"
	"time"
)

// FloodControl реализация интерфейса для проверки правил флуд-контроля.
type FloodControl struct {
	mu       sync.Mutex
	calls    map[int64][]time.Time // Карта для отслеживания вызовов метода Check для каждого пользователя
	interval time.Duration         // Интервал времени для проверки флуд-контроля
	maxCalls int                   // Максимальное количество вызовов в интервал времени
}

// NewFloodControl создает новый экземпляр FloodControl.
func NewFloodControl(interval time.Duration, maxCalls int) *FloodControl {
	return &FloodControl{
		calls:    make(map[int64][]time.Time),
		interval: interval,
		maxCalls: maxCalls,
	}
}

// Check проверяет флуд-контроль для указанного пользователя.
func (fc *FloodControl) Check(ctx context.Context, userID int64) (bool, error) {
	fc.mu.Lock()
	defer fc.mu.Unlock()

	now := time.Now()
	// Получаем список времен всех предыдущих вызовов для пользователя
	calls := fc.calls[userID]

	// Удаляем из списка времен вызовы, которые устарели и выходят за пределы интервала
	var validCalls []time.Time
	for _, callTime := range calls {
		if now.Sub(callTime) <= fc.interval {
			validCalls = append(validCalls, callTime)
		}
	}

	// Добавляем текущий вызов в список
	validCalls = append(validCalls, now)

	// Проверяем количество вызовов в интервале времени
	if len(validCalls) > fc.maxCalls {
		return false, nil // Превышен лимит вызовов
	}

	// Обновляем список вызовов для пользователя
	fc.calls[userID] = validCalls
	return true, nil // Флуд-контроль пройден
}
