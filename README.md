# fullstack test app
Static front end in `public/` (served from the edge) + a Go backend (Dockerfile) that owns `/api/*`.
Deploy as ngris type `fullstack`; Runtime: Port `8080`, Backend paths `/api`, Publish dir `public`.
Open `/` (static) then click the button → it fetches `/api/hello` from the backend pod.
