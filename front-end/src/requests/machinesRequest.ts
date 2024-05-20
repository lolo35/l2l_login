export const fetchMachines = async (url:string, lines: string) => {
    try {
        let headers = new Headers()
        headers.append("Content-Type", "application/x-www-form-urlencoded");

        let urlencoded = new URLSearchParams();
        urlencoded.append("lines", lines);

        const options = {
            method: "POST",
            headers: headers,
            body: urlencoded,
            //redirect: 'follow'
        };

        const request = await fetch(`${url}machines`, options)
        const response = await request.json();
        if(import.meta.env.VITE_MODE === "dev") console.log(response)
        return response
    } catch (exception) {
        if(exception instanceof Error) {
            if(import.meta.env.VITE_MODE === "dev") {
                throw new Error(exception.message)
            }
        }
    }
}