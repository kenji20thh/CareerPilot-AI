export type Job = { id: string; title: string; company: string; location: string; mode: string; salary: string; posted: string; match: number; tags: string[]; description: string }
export type Application = { id: string; title: string; company: string; status: 'Saved' | 'Applied' | 'Interview' | 'Offer'; date: string; match: number }
export type Profile = { name: string; headline: string; location: string; summary: string; skills: string[]; experience: { role: string; company: string; dates: string; bullets: string[] }[] }
export type Preferences = { titles: string; locations: string; remote: boolean; salary: string; industries: string }
export type ApiResult = { ok: boolean; status: number; data: unknown; error?: string }

export const emptyProfile: Profile = { name: '', headline: '', location: '', summary: '', skills: [], experience: [] }
export const defaultPreferences: Preferences = { titles: 'Product Designer, UX Designer', locations: 'New York, Remote', remote: true, salary: '$120k–$165k', industries: 'SaaS, Fintech, Consumer' } 
export const mockJobs: Job[] = [
  { id: 'j1', title: 'Senior Product Designer', company: 'Northstar Labs', location: 'New York, NY', mode: 'Hybrid', salary: '$140k–$175k', posted: '2d ago', match: 94, tags: ['Figma', 'Design systems', 'B2B SaaS'], description: 'Lead end-to-end product design for a small, ambitious team building the operating system for modern finance.' },
  { id: 'j2', title: 'Product Designer, Growth', company: 'Arc & Co.', location: 'Remote — US', mode: 'Remote', salary: '$125k–$155k', posted: '4d ago', match: 88, tags: ['Experimentation', 'Analytics', 'Growth'], description: 'Partner with growth, product, and engineering to turn customer insight into simple, high-converting experiences.' },
  { id: 'j3', title: 'UX Designer', company: 'Fieldwork', location: 'Brooklyn, NY', mode: 'On-site', salary: '$110k–$135k', posted: '1w ago', match: 79, tags: ['Research', 'Prototyping', 'Mobile'], description: 'Make complex tools feel clear and human for teams working in the field.' },
  { id: 'j4', title: 'Design Lead', company: 'Kindred Health', location: 'Remote — US', mode: 'Remote', salary: '$150k–$190k', posted: '1w ago', match: 74, tags: ['Leadership', 'Healthcare', '0→1'], description: 'Set the product design vision for a new patient experience from the ground up.' },
]
export const mockApplications: Application[] = [
  { id: 'a1', title: 'Senior Product Designer', company: 'Northstar Labs', status: 'Interview', date: 'Aug 12, 2026', match: 94 },
  { id: 'a2', title: 'Product Designer, Growth', company: 'Arc & Co.', status: 'Applied', date: 'Aug 10, 2026', match: 88 },
  { id: 'a3', title: 'Design Lead', company: 'Kindred Health', status: 'Saved', date: 'Aug 08, 2026', match: 74 },
]
export const mockProfile: Profile = { name: 'Maya Chen', headline: 'Product designer crafting calm, high-performing digital products', location: 'Brooklyn, NY', summary: 'Product designer with 7+ years turning ambiguous problems into clear, useful experiences. I care about systems, strong collaboration, and the small details that make products feel inevitable.', skills: ['Product strategy', 'UX research', 'Interaction design', 'Design systems', 'Figma', 'Prototyping'], experience: [{ role: 'Senior Product Designer', company: 'Lumen Finance', dates: '2022 — Present', bullets: ['Led redesign of core planning workflow used by 40k+ customers.', 'Built the first shared design system, reducing delivery time by 30%.'] }, { role: 'Product Designer', company: 'Studio Common', dates: '2019 — 2022', bullets: ['Shipped mobile and web experiences across fintech and healthcare.'] }] }
