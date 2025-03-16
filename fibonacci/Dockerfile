FROM node:18-alpine AS builder

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
COPY package*.json ./
COPY tsconfig.json ./
RUN npm install

# Bundle app source
COPY . .

# Build TypeScript code
RUN npm run build

# Production image
FROM node:18-alpine

WORKDIR /usr/src/app

# Copy package files and install only production dependencies
COPY package*.json ./
RUN npm install --only=production

# Copy built app from the builder stage
COPY --from=builder /usr/src/app/dist ./dist

# Expose port
EXPOSE 8080

# Start the application
CMD [ "node", "dist/index.js" ]