from http.server import HTTPServer, BaseHTTPRequestHandler

# In its base form it can't do anything 
# when you try to reach it , it gives you server error

httpd = HTTPServer(('localhost',8080), BaseHTTPRequestHandler)
httpd.serve_forever()