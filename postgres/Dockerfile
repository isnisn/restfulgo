# Start with an Alpine-based Node.js image
FROM node:18-alpine

# Install required packages including PostgreSQL client
RUN apk add --no-cache postgresql-client bash curl

# Install Newman globally
RUN npm install -g newman

# Set entrypoint for Newman
ENTRYPOINT ["newman"]
