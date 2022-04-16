package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	owm "github.com/briandowns/openweathermap"
)

func getWeather() int {
	apiKey := goDotEnvVariable("OWM_API_KEY")

	w, err := owm.NewCurrent("C", "en", apiKey)
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByName("Singapore")

	uv, err := owm.NewUV(apiKey)
	if err != nil {
		log.Fatalln(err)
	}

	if err := uv.Current(&w.GeoPos); err != nil {
		log.Fatalln(err)
	}

	info, err := uv.UVInformation()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\nToday's Weather:")
	fmt.Printf("UV Index: %.2f, Risk: %s\n", info[0].UVIndex[0], info[0].Risk)
	fmt.Println(info[0].RecommendedProtection)
	fmt.Printf("Temperature: %.2f°C\nFeels Like: %.2f°C\nMax Temp.: %.2f°C\nMin Temp.: %.2f°C\n", w.Main.Temp, w.Main.FeelsLike, w.Main.TempMax, w.Main.TempMin)
	fmt.Printf("Humidity: %d%%\n", w.Main.Humidity)
	fmt.Printf("Rain in the past hour: %.2fmm\n", w.Rain.OneH)
	fmt.Printf("Wind Speed: %.2f m/s\n", w.Wind.Speed)
	fmt.Printf("%s, Cloudiness: %d%%\n", w.Weather[0].Description, w.Clouds.All)
	getSuntimes(w.Sys)

	return w.Weather[0].ID
}

func getSuntimes(s owm.Sys) {
	dt := time.Now().Unix()
	// dt := int64(1649578993)
	dtSunrise := dt - int64(s.Sunrise)
	dtSunriseAbs := math.Abs(float64(dtSunrise))
	dtSunset := dt - int64(s.Sunset)
	dtSunsetAbs := math.Abs(float64(dtSunset))

	if dtSunrise <= 0 {
		if dtSunrise < -3600 {
			fmt.Println("Sunrise at", fmt.Sprint(time.Unix(int64(s.Sunrise), 0))[11:19])
		} else {
			dtSunriseTime, err := time.ParseDuration(strconv.Itoa(int(dtSunriseAbs)) + "s")
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("Sunrise in %s. Try looking out for it!\n", dtSunriseTime)
		}
		fmt.Println("Sunset at", fmt.Sprint(time.Unix(int64(s.Sunset), 0))[11:19])
	} else if dtSunset <= 0 {
		if dtSunset < -3600 {
			fmt.Println("Sunset at", fmt.Sprint(time.Unix(int64(s.Sunset), 0))[11:19])
		} else {
			dtSunsetTime, err := time.ParseDuration(strconv.Itoa(int(dtSunsetAbs)) + "s")
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("Sunset in %s. Try looking out for it!\n", dtSunsetTime)
		}
	}
}
