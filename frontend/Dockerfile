# Etapa de build
FROM node:24 AS builder

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

# Genera build de producción
RUN npm run build

# Etapa final: servir con nginx
FROM nginx:stable-alpine

# Copiar build de Vue al html de nginx
COPY --from=builder /app/dist /usr/share/nginx/html

# Copiar configuración personalizada de nginx (opcional)
# COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
