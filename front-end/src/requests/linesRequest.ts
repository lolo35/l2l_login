export const fetchLines = async (url:string, area_id:number) => {
    try {
        const request = await fetch(`${url}area_lines?area_id=${area_id}`);
        const response = await request.json();
        if(import.meta.env.VITE_MODE === "dev") console.log(`lines`, response);

        return response;
    } catch(exception) {
        if (exception instanceof Error) {
            if(import.meta.env.VITE_MODE === "dev") {
                throw new Error(exception.message);
            }
        }
    }
}