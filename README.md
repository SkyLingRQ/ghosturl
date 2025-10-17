# GhostURL
GhostURL es un programa de recopilación de URLs y subdominios históricos que Wayback Machine recopiló. Usa la API de CDX de Wayback Machine y ofrece escaneos por dominio único o fichero de dominios/subdominios. La facilitación de la obtención de URLs hacia un objetivo la hace una herramienta esencial, útil e indispensable en tareas de Read Team, Pentesting y Bug Bounty.
## Install
```bash
go install github.com/SkyLingRQ/ghosturl@latest
sudo mv go/bin/ghosturl /usr/bin
```
## Usage
```bash
Usage of GhostURL:
  -d string
    	Domain for the scan.
  -f string
    	File for the scan.
```
## Scan Single Domain
```bash
ghosturl -d example.com >> resultados.txt
```
## Scan Multiples Domains
```bash
ghosturl -f domains.txt >> resultados.txt
```
