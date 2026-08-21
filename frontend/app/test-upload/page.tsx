"use client";

import {useState} from "react";

interface uploadResponse {
    success : boolean;
    filename: string;
    message: string;

}

export default function TestUpload{
    const [file, setFile] = useState<File | null>(null);
    const [status, setStatus] = useState<string>("");

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        if (!file) {
            setStatus("Please select a file first")
            return;
        }

        const formData = new FormData 
        formData.append("cv", file)
        
    }
 }