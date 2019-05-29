var map = L.map('map', {
    attributionControl: false,
    boxZoom: false,
    doubleClickZoom: false,
    dragging: false,
    scrollWheelZoom: false,
    zoomSnap: 0,
});

map.addControl(L.control.attribution({
    prefix: "",
}))

var url ='https://api.tiles.mapbox.com/v4/{id}/{z}/{x}/{y}@2x.png?access_token={accessToken}';
var attribution = "Made with <a href='https://leafletjs.com/'>Leaflet</a>";
attribution += " | Map data © <a href='https://openstreetmap.org'>OpenStreetMap</a> contributors";
attribution += " | Imagery © <a href='https://www.mapbox.com/'>Mapbox</a>";

var imagery = new L.TileLayer(url, {
    accessToken: 'pk.eyJ1IjoianVsaWVuY2hlcnJ5IiwiYSI6ImNqbjhxM3F2czA0eWIza3J5YWY5YWQ3aG8ifQ.BiSRVINfh-aMDu4764_C7A',
    attribution: attribution,
    id: 'mapbox.streets',

    minZoom: 1,
    maxZoom: 13,
});

map.setView(new L.LatLng(42.3390559, -71.0897654), 13);
map.addLayer(imagery);
