function resolve() {
    let body = {
        "host": document.getElementById("host").value,
        "port": parseInt(document.getElementById("port").value),
        "user": document.getElementById("user").value,
        "password": document.getElementById("pass").value,
        "database": document.getElementById("db").value
    }

    fetch("/api/resolve", {
        method: "POST",
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(body)
    }).then(res => {
        if (res.status === 200) {
            alert("Database resolved!");
        } else {
            alert("Could not connect to database!");
        }
    }).catch(e => {
        console.error(e);
        alert("Gateway server error!")
    })
}
