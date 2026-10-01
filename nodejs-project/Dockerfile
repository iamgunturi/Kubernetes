# Use lightweight Node.js Alpine image
FROM node:18-alpine

# Set working directory inside container
WORKDIR /app

# Copy package.json and install dependencies
COPY package*.json ./
RUN npm install --production

# Copy project files
COPY . .

# Expose port
EXPOSE 3000

# Start the app
CMD ["npm", "start"]

