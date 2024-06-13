import React, {useEffect, useRef, useState} from 'react';
import './App.css';
// import {
//     YMaps,
//     Map,
//     Placemark,
//     Panorama,
//     useYMaps,
//     GeoObject,
//     FullscreenControl,
//     RulerControl, ZoomControl, Polygon
// } from '@pbe/react-yandex-maps';
// import {Map, View} from 'ol';
import Map from 'ol/Map';
import View from 'ol/View';
import TileLayer from 'ol/layer/Tile';
import TileWMS from "ol/source/TileWMS";
import OSM from "ol/source/OSM";
import "ol/ol.css";


function App() {
    return (
        <Mapp/>
    );
}


function Mapp() {
    useEffect(() => {
        const map = new Map({
            target: "map",
            layers: [
                new TileLayer({
                    source: new OSM(),
                }),
                // new TileLayer({
                //     source: new TileWMS({
                //         url: 'http://localhost:81/cgi-bin/myProject/qgis_mapserv.fcgi?SERVICE=WMS&VERSION=1.3.0&REQUEST=GetCapabilities',
                //         serverType: 'qgis',
                //         params: {},
                //     })
                // })
            ],
            view: new View({
                center: [51.505, -0.09],
                zoom: 5,
            }),
        });

        return () => {
            map.setTarget();
        };
    }, []);

    return (
        <div id="map" style={{width: "100%", height: "800px"}}/>
    )
}

export default App;






//
//
//     const [hint, setHint] = useState(null);
//
//     const handlePlacemarkMouseEnter = (e: any) => {
//         setHint(e.get('target').properties.get('hintContent'));
//     };
//
//     const handlePlacemarkMouseLeave = () => {
//         setHint(null);
//     };
//
//     return (
//         <Map defaultState={{ center: [66.25, 94.15], zoom: 4 }} style={{height: "80%", width: "80%"}} options={{
//             suppressMapOpenBlock: true, // Отключаем всплывающее окно с картой
//             yandexMapDisablePoiInteractivity: true, // Отключаем интерактивность POI
//         }}>
//             <FullscreenControl />
//             <RulerControl options={{ position: {right: 10, top: 50} }} />
//             <ZoomControl options={{ position: {right: 10, top: 100}, size: "small" }} />
//             <Placemark geometry={[66.25, 94.15]} options={{
//                 // iconColor: "red",
//                 iconLayout: 'default#image',
//                 // preset: "islands#circleDotIcon",
//                 iconImageHref: 'anchor.svg',
//                 iconImageSize: [20, 20],
//                 iconImageOffset: [-10, -10],
//                 hasHint: true,
//                 openHintOnHover: true
//             }}/>
//             <Polygon
//                 geometry={[
//                     // Координаты внешнего контура.
//                     [[55, 37], [56, 38], [67, 95], [66, 94]]
//                 ]}
//                 properties={{
//                     hintContent: "Многоугольник"
//                 }}
//                 options={{
//                     // Делаем полигон прозрачным для событий карты.
//                     interactivityModel: 'default#transparent',
//                     fillColor: '#F99',
//                     fillOpacity: 0.7,
//                     fillMethod: "tile",
//                     fillImageHref: 'anchor.svg',
//                     strokeColor: '#F55',
//                     strokeOpacity: 0.8,
//                     strokeWidth: 3,
//                 }}
//                 onMouseEnter={handlePlacemarkMouseEnter}
//                 onMouseLeave={handlePlacemarkMouseLeave}
//             />
//             <GeoObject
//                 geometry={{
//                     type: "LineString",
//                     coordinates: [
//                         [55.76, 37.64],
//                         [66.25, 94.15],
//                     ],
//                 }}
//                 properties={{
//                     hintContent: "Москва-Берлин"
//                 }}
//                 options={{
//                     geodesic: true,
//                     strokeWidth: 2,
//                     strokeColor: "#99F",
//                     strokeOpacity: 0.8
//                 }}
//                 onClick={handlePlacemarkMouseEnter}
//                 // onMouseEnter={handlePlacemarkMouseEnter}
//                 onMouseLeave={handlePlacemarkMouseLeave}
//             />
//             {hint && <div style={{ position: 'absolute', top: '50px', left: '50px', background: 'white', padding: '5px' }}>{hint}</div>}
//
//         </Map>
//     );
// }

