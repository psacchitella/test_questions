from http.server import BaseHTTPRequestHandler, HTTPServer

class HelloHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header("Content-type", "text/plain")
        self.end_headers()
        self.wfile.write(b"Hello, World!")

def run(server_class=HTTPServer, handler_class=HelloHandler):
    server_address = ('', 8080)
    httpd = server_class(server_address, handler_class)
    print("Server running at http://localhost:8080/")
    httpd.serve_forever()

if __name__ == "__main__":
    run()
