# CareerPilot — Project TODO

_Last updated: current session_

## ✅ Done

### Infrastructure
- [x] Go backend scaffolded (`main.go`, `handlers/`, `db/`, `middleware/`)
- [x] Next.js + TypeScript + Tailwind frontend scaffolded
- [x] CORS configured between frontend (`localhost:3000`) and backend (`localhost:8080`)
- [x] `.env` set up for secrets, `.gitignore`'d (rotated leaked Postgres password after accidental commit)

### Database
- [x] PostgreSQL installed and running locally
- [x] `careerpilot` database created
- [x] `pgcrypto` extension enabled (for UUID generation)
- [x] `users` table created (`id`, `email`, `password_hash`, `created_at`)
- [x] `cvs` table created (`id`, `user_id`, `filename`, `path`, `uploaded_at`)
- [x] Go connected to Postgres via `pgxpool` (`db.Pool`)

### Auth
- [x] `POST /api/signup` — validates input, hashes password with bcrypt, inserts user, returns JWT
- [x] `POST /api/login` — looks up user, verifies password with bcrypt, returns JWT
- [x] Generic "Invalid email or password" error on both bad email and bad password (prevents user enumeration)
- [x] JWT signing via `JWT_SECRET` env var
- [x] Auth middleware (`middleware.RequireAuth`) — verifies `Authorization: Bearer <token>` header, injects `user_id` into request context

### File Upload
- [x] `POST /api/upload-cv` — accepts multipart file upload
- [x] File size capped at 5MB
- [x] Saves file to `uploads/` directory on disk
- [x] Now requires auth (wrapped in `RequireAuth` middleware)
- [x] Inserts a row into `cvs` linking `user_id`, `filename`, `path`
- [x] Tested end-to-end from the Next.js frontend (`/test`, `/test-upload` pages)

### Endpoints summary
| Method | Route | Auth required | Status |
|---|---|---|---|
| GET | `/api/health` | No | ✅ Working |
| POST | `/api/signup` | No | ✅ Working |
| POST | `/api/login` | No | ✅ Working |
| POST | `/api/upload-cv` | Yes | ⚠️ Needs re-test after middleware wiring |

---

## 🔲 Immediate next steps

- [ ] **Re-test `/api/upload-cv` end-to-end** with the new auth middleware:
  - [ ] Confirm it returns `401 Unauthorized` with no token
  - [ ] Confirm it returns `401` with an invalid/expired token
  - [ ] Confirm it succeeds with a valid token from `/api/login`
  - [ ] Confirm a row actually lands in the `cvs` table with the correct `user_id`
- [ ] Update the frontend upload form/page to:
  - [ ] Store the JWT after login (e.g. in memory / React state — not `localStorage` in artifacts, but fine in a real Next.js app)
  - [ ] Send it as `Authorization: Bearer <token>` on the upload request
- [ ] Build simple login/signup pages in the frontend (currently only tested via curl)

## 🔲 Hardening (uploads)

- [ ] File-type validation — restrict to `.pdf`, `.doc`, `.docx` (currently accepts anything)
- [ ] Filename collision handling — two users uploading `resume.pdf` currently overwrite each other; prefix with UUID or timestamp
- [ ] Move file storage off local disk before any real deployment (e.g. S3/GCS/R2) — local disk won't survive redeploys on most hosting platforms

## 🔲 Auth hardening / completeness

- [ ] `GET /api/me` — return the logged-in user's info given a valid token (useful for frontend to check "am I logged in")
- [ ] Password reset flow (email-based, later)
- [ ] Rate limiting on `/api/login` and `/api/signup` (prevent brute-force / spam signups)
- [ ] Input validation — proper email format check, minimum password length/strength on signup

## 🔲 General backend

- [ ] Replace `println`/`panic` with structured logging (`log/slog` or similar)
- [ ] Centralized error-response helper (consistent JSON error shape across all endpoints, not just `http.Error` plain text)
- [ ] Environment-based config (dev vs. prod `DATABASE_URL`, CORS origin, etc.)
- [ ] Basic request logging middleware (method, path, status, latency)

## 🔲 Core CareerPilot product features (not started)

This is the actual product — everything above is foundation/infrastructure.

- [ ] CV parsing (extract text/structured data from uploaded PDF/DOCX)
- [ ] Whatever the core AI feature is meant to be (e.g. job matching, CV scoring/feedback, cover letter generation — needs definition)
- [ ] User dashboard (view uploaded CVs, past results, etc.)
- [ ] Job listing storage/ingestion (if matching against real listings)

## 🔲 Deployment (later)

- [ ] Choose hosting (backend: Fly.io / Railway / Render; frontend: Vercel is the natural fit for Next.js)
- [ ] Managed Postgres (e.g. Supabase, Neon, Railway) instead of local
- [ ] Environment variables set in hosting platform (never commit secrets again)
- [ ] HTTPS / real domain
- [ ] Restrict CORS to the real production frontend origin (currently hardcoded to `localhost:3000`)

---

## Notes / lessons from this session
- Windows PowerShell's `curl` is aliased to `Invoke-WebRequest`, not real curl — use `curl.exe` explicitly, and remember `-d "@file.json"` needs the `@` to read from a file.
- `Out-File -Encoding utf8` in PowerShell adds a BOM that breaks JSON parsing — use `[System.IO.File]::WriteAllText(...)` instead for clean JSON test files.
- Go doesn't hot-reload — always restart (`Ctrl+C` then `go run main.go`) after any code or `.env` change.
- Never commit `.env` — if it happens, rotate the leaked credentials immediately, not just remove the file from tracking.