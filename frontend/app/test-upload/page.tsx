"use client";

import { useState } from "react";

interface UploadResponse {
  success: boolean;
  filename: string;
  message: string;
}

export default function TestUpload() {
  const [file, setFile] = useState<File | null>(null);
  const [status, setStatus] = useState<string>("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!file) {
      setStatus("Please select a file first.");
      return;
    }

    const formData = new FormData();
    formData.append("cv", file);

    try {
      const res = await fetch("http://localhost:8080/api/upload-cv", {
        method: "POST",
        body: formData,
      });

      if (!res.ok) {
        const errText = await res.text();
        setStatus(`Error: ${errText}`);
        return;
      }

      const data: UploadResponse = await res.json();
      setStatus(data.message);
    } catch (err) {
      console.error(err);
      setStatus("Upload failed — check console.");
    }
  };

  return (
    <form onSubmit={handleSubmit} className="flex flex-col gap-4 max-w-sm p-6">
      <input
        type="file"
        accept=".pdf,.doc,.docx"
        onChange={(e) => setFile(e.target.files?.[0] ?? null)}
        className="border border-gray-300 rounded-md p-2 text-sm"
      />
      <button
        type="submit"
        className="bg-blue-600 text-white rounded-md py-2 px-4 hover:bg-blue-700 transition"
      >
        Upload CV
      </button>
      {status && <p className="text-sm text-gray-700">{status}</p>}
    </form>
  );
}