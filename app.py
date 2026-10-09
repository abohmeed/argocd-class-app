from http.server import BaseHTTPRequestHandler, HTTPServer

VERSION = "v1"


class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.end_headers()
        self.wfile.write(f"storefront {VERSION}".encode())


if __name__ == "__main__":
    HTTPServer(("0.0.0.0", 5678), Handler).serve_forever()
