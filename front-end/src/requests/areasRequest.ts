export const fetchAreas = async (url:string) => {
    try {
        const request = await fetch(`${url}areas`);
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