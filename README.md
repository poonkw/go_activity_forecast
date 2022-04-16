Add a .env file with OWM_API_KEY after creating one via OpenWeatherMap

Open Weather Documentation: 
- https://github.com/briandowns/openweathermap
- https://pkg.go.dev/github.com/briandowns/openweathermap
- https://openweathermap.org/current

Run:
```
go run .\main.go .\weather.go .\getenv.go
```

Sample Output:
```
Welcome to your day planner. Please select the following:
1) Show me what I can do today!
2) Add an interesting activity
3) Remove an activity for the foreseeable future
1

Today's Weather:
UV Index: 11.00, Risk: Extreme
Take all precautions: Wear SPF 30+ sunscreen, a long-sleeved shirt and trousers, sunglasses, and a very broad hat. Avoid the sun within three hours of solar noon.
Temperature: 31.71°C
Feels Like: 37.29°C
Max Temp.: 32.97°C
Min Temp.: 30.02°C
Humidity: 63%
Rain in the past hour: 0.00mm
Wind Speed: 5.66 m/s
broken clouds, Cloudiness: 75%
Sunset at 19:08:34

Recommended activites:
Sunbathing
Reading
Cycling
Yoga
```
