## Polecenie użyte do zbudowania i wypchnięcia obrazu

```bash
docker buildx build \
  --builder testbuilder \
  --platform linux/amd64,linux/arm64 \
  --tag mikitax/weather-app-dot:latest \
  --cache-to type=registry,ref=mikitax/weather-app-dot:buildcache,mode=max \
  --cache-from type=registry,ref=mikitax/weather-app-dot:buildcache \
  --ssh default \
  --push .
```

## Potwierdzenie architektur (Manifest)

```bash
docker buildx imagetools inspect mikitax/weather-app-dot:latest
```

<img width="940" height="383" alt="image" src="https://github.com/user-attachments/assets/3fb44845-1b9a-4837-a0ca-011ded23cc0a" />

## Dowód utworzenia obrazu cache

```bash
docker buildx imagetools inspect mikitax/weather-app-dot:buildcache
```

<img width="940" height="96" alt="image" src="https://github.com/user-attachments/assets/98663001-f9cc-481b-a54f-3817ed3c248d" />

## Dowód poprawnego wykorzystywania danych cache

<img width="940" height="497" alt="image" src="https://github.com/user-attachments/assets/253c4fe2-1588-4edc-9e68-c8cdfdb59db0" />

## Skanowanie CVE

```bash
docker scout cves mikitax/weather-app-dot:latest
```

<img width="1054" height="450" alt="image" src="https://github.com/user-attachments/assets/e5c381e3-a172-44ec-b75f-278b700170c9" />

##Wygląd aplikacji

<img width="897" height="592" alt="image" src="https://github.com/user-attachments/assets/6c136c6a-b159-4e2b-a50a-cc79d67b207d" />


