# MpangoWaCuddles.com

Curated discovery platform for budget-friendly date spots around Nairobi. This repo contains:

- Go (Gin) backend scaffold with demo data and REST endpoints.
- Next.js frontend (app router) with romantic, mobile-first UI.

## Project layout

```
mpango-wa-cuddles/
├── cmd/server/main.go         # Go entrypoint
├── internal/                  # Handlers, models, router
└── web/nextjs-app/            # Next.js app router UI
```

## Quickstart: backend (demo mode)

1) Ensure Go 1.22+ is installed.  
2) Run the API:

```bash
go run ./cmd/server
```

Endpoints (mock/in-memory):

- `GET /healthz`
- `GET /locations` with filters `?cost=200-500&area=Karen&activity=Picnic`
- `GET /locations/:id`
- `GET /ads`
- `POST /tickets/initiate` `{ "location_id": "...", "phone": "07..." }`
- `GET /tickets/status/:id`

Admin stubs exist for locations and ads (replace with DB + auth).

## Quickstart: frontend

```bash
cd web/nextjs-app
npm install
npm run dev
```

Open http://localhost:3000 to view the mobile-first UI. The UI currently reads demo data; wire it to the Go API via `fetch` to go live.

## Next implementation steps

- Replace in-memory demo data with PostgreSQL using GORM/pgx.
- Add JWT auth and session-backed CuddleList.
- Implement M-Pesa STK push flow and callback handler.
- Build admin UI for adding locations/ads and uploading images to S3.
- Add weighted ad rotation + click tracking.
- Deploy: backend to Render/Railway, frontend to Vercel/Netlify.


