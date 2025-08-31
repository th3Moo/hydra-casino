# Use official Nginx base image
FROM nginx:1.25-alpine

# Copy custom nginx.conf
COPY nginx.conf /etc/nginx/nginx.conf

# Copy frontend static files into the default nginx web root
COPY ../frontend /usr/share/nginx/html

# Expose port 80 for HTTP traffic
EXPOSE 80

# Start nginx in the foreground
CMD ["nginx", "-g", "daemon off;"]
