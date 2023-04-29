from http.server import HTTPServer, BaseHTTPRequestHandler

class myServer(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == '/':
            self.path = '/index.html'
        try:
            file_to_open = open('static/'+self.path[1:]).read()
            # file_to_open = f"File  Found {'static/'+self.path[1:]}"
            self.send_response(200)
        except:
            file_to_open = f"File Not Found {'static/'+self.path[1:]}"
            self.send_response(404)
        self.end_headers()
        self.wfile.write(bytes(file_to_open,'utf-8'))
        
    def do_POST(self):
            self.send_response(200)
            self.send_header('Content-type', 'text/html')
            self.end_headers()
            response = 'Received POST data'
            self.wfile.write(response.encode('utf-8'))

httpd = HTTPServer(('',8080), myServer)
httpd.serve_forever()