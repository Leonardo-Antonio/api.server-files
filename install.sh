echo "#####################Acualizando repositorios#####################"
sudo apt update -y

echo "#####################Instalación de curl y git#####################"
sudo apt install curl git -y

echo "#####################Paquetes de requisitos previos#####################"
sudo apt install apt-transport-https ca-certificates curl software-properties-common -y

echo "#####################Clave de GPG para el repositorio oficial de Docker#####################"
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

echo "#####################Repositorio de Docker a las fuentes de APT#####################"
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu focal stable" -y

echo "#####################Acualizando repositorios#####################"
sudo apt update -y

apt-cache policy docker-ce

echo "#####################Instalando Docker#####################"
sudo apt install docker-ce -y

echo "#####################Estableciando que el sistema siempre inicie con los servicios de docker en 'on'#####################"
sudo systemctl enable docker

echo "#####################Estado del servicio de docker docker#####################"
sudo systemctl status docker

echo "#####################Instalación de Docker Compose#####################"
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

echo "#####################Añadiendo permisos#####################"
sudo chmod +x /usr/local/bin/docker-compose
