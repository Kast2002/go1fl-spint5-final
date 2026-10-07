package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("Некорректное число шагов")
	}
	if weight < 20 || weight > 300 {
		return 0, errors.New("Некорректный вес")
	}
	if height < 0.5 || height > 3 {
		return 0, errors.New("Некорректный рост")
	}
	if duration <= 0 {
		return 0, errors.New("Некорректное время")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	return weight * meanSpeed * duration.Minutes() * walkingCaloriesCoefficient / minInH, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("Некорректное число шагов")
	}
	if weight < 20 || weight > 300 {
		return 0, errors.New("Некорректный вес")
	}
	if height < 0.5 || height > 3 {
		return 0, errors.New("Некорректный рост")
	}
	if duration <= 0 {
		return 0, errors.New("Некорректное время")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	return weight * meanSpeed * duration.Minutes() / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	if steps <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 {
		return 0
	}
	lengthStep := height * stepLengthCoefficient
	return lengthStep * float64(steps) / mInKm
}
