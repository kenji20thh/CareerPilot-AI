"use client";

import { useEffect } from "react";

export default function TestPage() {
  useEffect(() => {
    fetch("http://localhost:8080/api/health")
      .then((res) => res.json())
      .then((data: { status: string }) => console.log(data))
      .catch((err) => console.error("Fetch failed:", err));
  }, []);

  return <div className="p-6 text-lg">Check the console for backend health status.</div>;
}