import type { ApiResult } from './types'

const base = process.env.NEXT_PUBLIC_API_URL?.replace(/\/$/, '')
export async function apiFetch(path: string, init?: RequestInit): Promise<ApiResult> {
  if (!base) return { ok: true, status: 200, data: null }
  try {
    const response = await fetch(`${base}${path}`, { ...init, headers: { ...(init?.body instanceof FormData ? {} : { 'Content-Type': 'application/json' }), ...init?.headers } })
    const data = await response.json().catch(() => null)
    return { ok: response.ok, status: response.status, data, error: response.ok ? undefined : data?.message || 'Request failed' }
  } catch (error) { return { ok: false, status: 0, data: null, error: error instanceof Error ? error.message : 'Network error' } }
}
