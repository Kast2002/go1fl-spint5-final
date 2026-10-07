package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	personaldata.Personal
	Steps    int
	Duration time.Duration
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	dataSlice := strings.Split(datastring, ",")
	if len(dataSlice) != 2 {
		return errors.New("Некорректный формат данных")
	}

	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("Некорректное число шагов")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(dataSlice[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("Некорректное время")
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf(`Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
`,
		ds.Steps, distance, calories)
	return result, nil
}
