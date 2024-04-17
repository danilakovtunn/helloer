all: minikube/install kubectl/install minikube/start deploy_all get_ip

minikube/install:
	curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
	sudo install minikube-linux-amd64 /usr/local/bin/minikube && rm minikube-linux-amd64

minikube/uninstall:
	sudo rm -rf /usr/local/bin/minikube
	rm -rf ~/.minikube

minikube/start:
	minikube start

minikube/stop:
	minikube stop

minikube/delete:
	minikube delete

get_ip:
	kubectl get nodes -o wide

kubectl/install:
	curl -LO https://dl.k8s.io/release/`curl -LS https://dl.k8s.io/release/stable.txt`/bin/linux/amd64/kubectl
	chmod +x ./kubectl
	sudo mv ./kubectl /usr/local/bin/kubectl

kubectl/uninstall:
	sudo rm -rf /usr/local/bin/kubectl

deploy_all:
	kubectl apply -f "./kube_files/*.yaml"

remove_all:
	kubectl delete -f "./kube_files/*.yaml"