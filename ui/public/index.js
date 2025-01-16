const hashHolder = document.getElementById("hash")
const plainTextData = document.getElementById("plain-text")

if(!hashHolder || !plainTextData){
    throw("the needed elements are missing.")
}

plainTextData.addEventListener("keyup", async (e)=>{
    if(e.key == 'Enter'){
        const pt = plainTextData.value
        // fetch here
        const response = await fetch("http://localhost:8080/enc", {
            method: "POST",
            body: pt
        })

        const sha256 = await response.text()
        hashHolder.innerHTML = sha256
    }
})

