# ngris-test-apps

Test applications for validating **git-source** deploys on ngris. ngris builds a git app
from the **repo root**, so each app lives on its own branch (with the app at the root);
`main` collects all three in subdirectories for easy browsing.

| App          | Branch                              | ngris type  | Exercises |
|--------------|-------------------------------------|-------------|-----------|
| Static site  | [`static`](../../tree/static)       | `static`    | git → build/publish → edge serve (+ custom 404) |
| Go backend   | [`backend`](../../tree/backend)     | `managed`   | git → signed/scanned image → run in-cluster → serve |
| Full-stack   | [`fullstack`](../../tree/fullstack) | `fullstack` | dual artifact: static front end on the edge **and** a backend image; `/api/*` → pod, everything else → edge |

## Deploy (Connect Git → this repo → the branch above)
- **static** — nothing extra.
- **backend** — Runtime: Port `8080`, Backend paths `/`.
- **fullstack** — Runtime: Port `8080`, Backend paths `/api`, Publish dir `public`.
