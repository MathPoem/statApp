package pkg

import (
	"fmt"
	"strconv"
	"strings"
)

func Stat(info *ResponseCurrencyInfo) *ResponseStat {
	stat := ResponseStat{
		Sum:          0,
		SumMonths:    make([]*SumMonth, 0, 12),
		DayPriceAvg:  make([]*GasAvg, 0, 365),
		Distribution: make([]*HourDistribution, 24),
	}

	countTransactionPerDay := make([]int, 365)

	for i := 0; i < 24; i++ {
		stat.Distribution[i] = &HourDistribution{}
		stat.Distribution[i].Hour = fmt.Sprintf("%02d:00", i+1)
		stat.Distribution[i].Statistic = make([]*HourDistribution_TimeStat, 20)
		for j := 0; j < 20; j++ {
			interval, _ := Interval(float64(j*10) + 19)
			stat.Distribution[i].Statistic[j] = &HourDistribution_TimeStat{
				Price: interval,
				Count: 0,
			}
		}
	}

	for _, v := range info.GetEthereum().GetTransactions() {
		t := strings.Split(v.GetTime(), " ")
		date := t[0]
		hour, _ := strconv.Atoi(t[1][:2])
		month, _ := strconv.Atoi(date[3:5])
		day, _ := strconv.Atoi(date[6:])

		// calculate common sum
		stat.Sum += v.GetGasPrice() * v.GetGasValue()

		// calculate sum gas per month
		if month > len(stat.SumMonths) {
			stat.SumMonths = append(stat.SumMonths, &SumMonth{
				Month:  date[:5],
				SumGas: v.GetGasValue(),
			})
		} else {
			stat.SumMonths[month-1].SumGas += v.GetGasValue()
		}

		// calculate distribution price
		_, indexInterval := Interval(v.GetGasPrice())
		stat.Distribution[hour].Statistic[indexInterval].Count += 1

		// calculate sum gas prices per day
		if monthToDay(month)+day > len(stat.DayPriceAvg) {
			stat.DayPriceAvg = append(stat.DayPriceAvg, &GasAvg{
				Time:        date,
				GasPriceAvg: v.GetMaxGasPrice(),
			})
		} else {
			stat.DayPriceAvg[monthToDay(month)+day-1].GasPriceAvg += v.GetGasPrice()
		}
		countTransactionPerDay[monthToDay(month)+day-1] += 1
	}

	// division sum gas prices on count transactions
	for i, v := range stat.DayPriceAvg {
		v.GasPriceAvg /= float64(countTransactionPerDay[i])
	}

	return &stat
}

func monthToDay(month int) int {
	count := 0
	if month == 1 {
		return count
	}
	count += 31
	if month == 2 {
		return count
	}
	count += 28
	if month == 3 {
		return count
	}
	count += 31
	if month == 4 {
		return count
	}
	count += 30
	if month == 5 {
		return count
	}
	count += 31
	if month == 6 {
		return count
	}
	count += 30
	if month == 7 {
		return count
	}
	count += 31
	if month == 8 {
		return count
	}
	count += 31
	if month == 9 {
		return count
	}
	count += 30
	if month == 10 {
		return count
	}
	count += 31
	if month == 11 {
		return count
	}
	count += 30
	return count
}

func Interval(val float64) (string, int) {
	if val < 20 {
		return "0:20", 0
	} else if val < 30 {
		return "20:30", 1
	} else if val < 40 {
		return "30:40", 2
	} else if val < 50 {
		return "40:50", 3
	} else if val < 60 {
		return "50:60", 4
	} else if val < 70 {
		return "60:70", 5
	} else if val < 80 {
		return "70:80", 6
	} else if val < 90 {
		return "80:90", 7
	} else if val < 100 {
		return "90:100", 8
	} else if val < 110 {
		return "100:110", 9
	} else if val < 120 {
		return "110:120", 10
	} else if val < 130 {
		return "120:130", 11
	} else if val < 140 {
		return "130:140", 12
	} else if val < 150 {
		return "140:150", 13
	} else if val < 160 {
		return "150:160", 14
	} else if val < 170 {
		return "160:170", 15
	} else if val < 180 {
		return "170:180", 16
	} else if val < 190 {
		return "180:190", 17
	} else if val < 200 {
		return "190:200", 18
	} else {
		return "200:", 19
	}
}
