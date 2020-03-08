# veterejo
Veterejo is a Go package for querying weathe data from Open Weather Maps.

Vetero is "weather" in Esperanto and "ejo" is place.
By replacing the **O** of a noun with **ejo** the word discribes a place of that thing.
(eg: Computilejo = computer-place or in English a computer lab)

## Using This Package:
1. Import
2. Create a new `WeatherData` struct: `w := new(veterejo.WeatherData)`
3. Call the data: `veterejo.w.getData(yourCityID, unitsOfMeasure, yourApiId)`
4. Use the method for temperature to get the most recent temperature data on OpenWeatherMap

*veterejo is licensed under the BSD 3-Clause "New" or "Revised" License*
