#  Demo: Resiliencia, Balanceo de Carga y Auto-recuperación

Este proyecto es un laboratorio práctico diseñado para demostrar cómo un orquestador (Docker) y un balanceador de carga (HAProxy) trabajan juntos para mantener un sistema en línea, incluso cuando los servidores fallan internamente.

##  Requisitos Previos

Solo necesitas una cosa para ejecutar este proyecto:
- **Docker Desktop** (en Windows/Mac) o **Docker Compose** (en Linux).
- *Nota: No es necesario tener instalado Go, ya que la compilación se realiza de forma aislada dentro de los contenedores.*

##  Cómo iniciar el proyecto

Abre una terminal en la carpeta raíz del proyecto y ejecuta:

```bash
docker-compose up --build
```

Este comando construirá las imágenes, limpiará configuraciones y levantará la infraestructura: un balanceador de carga y tres instancias de nuestra API.

## Rutas Importantes

Una vez que la terminal muestre que los servidores están escuchando, puedes usar las siguientes rutas en tu navegador:

- **API Principal:** [http://localhost](http://localhost) 
  *(Muestra qué servidor está respondiendo a la petición).*
- **Dashboard de HAProxy:** [http://localhost:8404/stats](http://localhost:8404/stats) 
  *(Panel visual en tiempo real para ver la salud de los nodos).*
- **Endpoint de Fallo (Crash):** [http://localhost/crash](http://localhost/crash) 
  *(Ruta "trampa" que fuerza el apagado repentino del servidor que la recibe).*

---


### Iniciar el tráfico constante
Para no tener que actualizar el navegador manualmente, abre una **nueva terminal** y ejecuta uno de los siguientes scripts para simular usuarios haciendo peticiones cada medio segundo:

**Si usas Ubuntu, Mac o Git Bash (Windows):**
```bash
while true; do curl -s http://localhost; sleep 0.5; done
```

**Si usas PowerShell (Windows):**
```powershell
while ($true) { curl.exe -s http://localhost; Start-Sleep -Milliseconds 500 }
```
*Notarás en la terminal cómo el balanceador reparte el tráfico equitativamente entre el Servidor 1, 2 y 3 (Round-Robin).*

### Mostrar la salud del sistema
Abre en tu navegador el **Dashboard de HAProxy** (`http://localhost:8404/stats`). En la sección `api_servers`, verás los tres nodos en color **verde** (Saludables). HAProxy hace un "Health Check" constante para verificar esto.

### Simular la caída
Abre una nueva pestaña en el navegador y entra a `http://localhost/crash`. 
Esto apagará inmediatamente uno de los servidores.

### Observar la Resiliencia y Auto-recuperación
Vuelve rápidamente al Dashboard y a tu terminal con el script:
1. **HAProxy** detectará que el nodo no responde, lo pondrá en **rojo** y dejará de enviarle tráfico. Los usuarios (tu script de la terminal) seguirán recibiendo respuestas de los otros dos servidores supervivientes sin notar interrupciones.
2. **Docker** detectará que el proceso murió. Gracias a la política `restart: always`, esperará 5 segundos (simulando un arranque lento) y revivirá el contenedor.
3. El nodo en el Dashboard volverá a color **verde** y se reintegrará automáticamente a la rotación de tráfico en tu terminal.
