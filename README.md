# veterejo
Veterejo is a Go package for querying weathe data from Open Weather Maps.

Vetero is "weather" in Esperanto and "ejo" is place.
By replacing the **o** of a noun with **ejo** the word discribes a place of that thing.
(eg: Computilejo = computer-place or in English a computer lab)

## Using This Package:
1. Get an API Key from [OpenWeatheMap](https://openweathermap.org)
2. Import this package
3. Create a new `WeatherData` struct: `w := new(veterejo.WeatherData)`
4. Parse the URL: `url, err := veterejo.MakeURL(yourCityID, unitsOfMeasure, yourApiId)`
   - `yourApiId` is required
   - The rest may be omitted if desired.
5. Call the data: `w.getData(url)`
6. Use the appropriate method to get the most recent data needed from OpenWeatherMap.

*veterejo is licensed under the BSD 3-Clause "New" or "Revised" License*
