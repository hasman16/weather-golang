# weather
Weather forecast using Golang and React

Description:

The API gets called from the React frontend, passing the latitude, longitude and datetime to the Get url. The Golang API receives the params and build the Dark Sky url on the fly and make a Restfull call to darksky.net which returns a payload of weather data.
The Golang API therefore returns that payload to the React frontend that initiated the call.
