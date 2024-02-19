
const getbaseURI = () => { 
    const env = process?.env?.NODE_ENV;
    if (env === 'prod'){
        return "http://65.0.129.57:8000"
    }else{
        return 'http://localhost:8000'
    }
}
export {
    getbaseURI
}