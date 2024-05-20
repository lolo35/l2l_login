export const login = async (url:string, lineids:string, machineids:string, operatorid:string) => {
    try {
        let headers = new Headers();
        headers.append("Content-Type", "application/x-www-form-urlencoded");

        let urlencoded = new URLSearchParams();
        urlencoded.append("lines", lineids);
        urlencoded.append("machines", machineids);
        urlencoded.append("operator", operatorid)

        const options = {
            method: "POST",
            headers: headers,
            body: urlencoded
        };

        const request = await fetch(`${url}login`, options);
        const response = await request.json();
        if(import.meta.env.VITE_MODE === "dev") console.log(response)
        return response;
    } catch (exception) {
        if(exception instanceof Error) {
            if(import.meta.env.VITE_MODE === "dev") {
                throw new Error(exception.message)
            }
        }
    }
}