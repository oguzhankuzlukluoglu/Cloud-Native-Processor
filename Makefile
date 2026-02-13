deploy:
	docker build -t cloud-native-api:v1 ./api
	docker build -t cloud-native-worker:v1 ./worker
	kubectl apply -f k8s/

clean:
	kubectl delete -f k8s/

status:
	kubectl get pods,svc,pvc