# Mars Rover Photos
This application shows mars rover photos searched by earth date.

# About NASA API

## Mars Rover Photos
This API is designed to collect image data gathered by NASA's Curiosity, Opportunity, and Spirit rovers on Mars and make it more easily available to other developers, educators, and citizen scientists. This API is maintained by Chris Cerami.

Each rover has its own set of photos stored in the database, which can be queried separately. There are several possible queries that can be made against the API. Photos are organized by the sol (Martian rotation or day) on which they were taken, counting up from the rover's landing date. A photo taken on Curiosity's 1000th Martian sol exploring Mars, for example, will have a sol attribute of 1000. If instead you prefer to search by the Earth date on which a photo was taken, you can do that too.

Along with querying by date, results can also be filtered by the camera with which it was taken and responses will be limited to 25 photos per call. Queries that should return more than 25 photos will be split onto several pages, which can be accessed by adding a 'page' param to the query.

Each camera has a unique function and perspective, and they are named as follows:

### Rover Cameras
| Abbreviation | Camera                                             | Curiosity | Opportunity | Spirit |
|:-------------|:---------------------------------------------------|:--------- |:------------|:-------|
| FHAZ         | Front Hazard Avoidance Camera                      | ✔         | ✔           | ✔      |
| RHAZ         | Rear Hazard Avoidance Camera                       | ✔         | ✔           | ✔      |
| MAST         | Mast Camera                                        | ✔         |             |        |
| CHEMCAM      | Chemistry and Camera Complex                       | ✔         |             |        |
| MAHLI        | Mars Hand Lens Imager                              | ✔         |             |        |
| MARDI        | Mars Descent Imager                                | ✔         |             |        |
| NAVCAM       | Navigation Camera                                  | ✔         | ✔           | ✔      |
| PANCAM       | Panoramic Camera                                   |           | ✔           | ✔      |
| MINITES      | Miniature Thermal Emission Spectrometer (Mini-TES) |           | ✔           | ✔      |

### Querying by Earth date
| Parameter  | Type       | Default  | Description                                   |
|:-----------|:-----------|:-------- |:----------------------------------------------|
| earth_date | YYYY-MM-DD | none     | corresponding date on earth for the given sol |
| camera     | string     | all      | see table above for abbreviations             |
| page       | int        | 1        | 25 items per page returned                    |
| api_key    | string     | DEMO_KEY | api.nasa.gov key for expanded usage           |

#### Example query
https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?earth_date=2015-6-3&api_key=DEMO_KEY

# How to test

## Run with Docker
```sh
git clone https://github.com/bigwindev/mars-rover-photos.git
cd mars-rover-photos
docker-compose up --remove-orphans
```

## Run without Docker
```sh
git clone https://github.com/bigwindev/mars-rover-photos.git
cd mars-rover-photos
go build
./mars-rover-photos
```

## Request on Web browser
```
http://localhost:8080
```
