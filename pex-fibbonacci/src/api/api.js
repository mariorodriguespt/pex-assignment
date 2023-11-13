const BASE_URL = "http://127.0.0.1:8080";
const globalOptions = {};

export const checkForExistingVisitor = async () => {
    const savedVisitor = localStorage.getItem("visitor-id");

    if(savedVisitor){
        globalOptions.headers =  { 'Authorization' : `Bearer ${savedVisitor}`};
    }
}

export const getCurrentNumber = async () => {    
    const response = await fetch(`${BASE_URL}/current`, globalOptions);
    const data = await response.json();        

    if(data.visitorId) {
        globalOptions.headers =  { 'Authorization' : `Bearer ${data.visitorId}`};
        localStorage.setItem("visitor-id", data.visitorId);
    }

    return data.number;    
};

export const getNextNumber = async () => {    
    const response = await fetch(`${BASE_URL}/next`, {
        method: "post",
        ...globalOptions
    });        
    const data = await response.json();        
    
    return data.number;   
};

export const getPreviousNumber = async () => {    
    const response = await fetch(`${BASE_URL}/previous`, {
        method: "post",
        ...globalOptions
    });
    const data = await response.json();        
    return data.number;    
};
