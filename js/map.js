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

var url ='https://api.mapbox.com/styles/v1/{id}/tiles/{z}/{x}/{y}?access_token={accessToken}';
var attribution = "Made with <a href='https://leafletjs.com/'>Leaflet</a>";
attribution += " | Map data © <a href='https://openstreetmap.org'>OpenStreetMap</a> contributors";
attribution += " | Imagery © <a href='https://www.mapbox.com/'>Mapbox</a>";

var imagery = new L.TileLayer(url, {
    accessToken: 'pk.eyJ1IjoianVsaWVuY2hlcnJ5IiwiYSI6ImNqbjhxM3F2czA0eWIza3J5YWY5YWQ3aG8ifQ.BiSRVINfh-aMDu4764_C7A',
    attribution: attribution,
    detectRetina: true,
    id: 'mapbox/streets-v11',
    maxZoom: 14,
    minZoom: 1,
    tileSize: 512,
    zoomOffset: -1,
});

map.setView(new L.LatLng(40.7080556, -73.9141667), 12);
map.addLayer(imagery);
