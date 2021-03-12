# relaxdays-hackathon-vol1-backend

This project was created in the Relaxdays Code Challenge Vol. 1. See https://sites.google.com/relaxdays.de/hackathon-relaxdays/startseite for more information. My participant ID in the challenge was: CC-VOL1-7

## How to run this project

You can get a running version of this code by using:

```bash
git clone https://github.com/hashworks/relaxdays-hackathon-vol1-backend.git
cd relaxdays-hackathon-vol1-backend
docker build -t relaxdays-hackathon-vol1-backend .
docker run -p 8080:8080 -it relaxdays-hackathon-vol1-backend
```

Afterwards you can access http://127.0.0.1:8080/ which will redirect you to the swagger UI.