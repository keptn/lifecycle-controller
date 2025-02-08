LFC_NAMESPACE ?= keptn-system
PODTATO_NAMESPACE ?= podtato-kubectl
ARGO_NAMESPACE ?= argocd
# renovate: datasource=github-tags depName=argoproj/argo-cd
ARGO_VERSION ?= v2.13.5
ARGO_SECRET = $(shell kubectl -n ${ARGO_NAMESPACE} get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo)

.PHONY: install
install:
	@echo "-----------------------------------"
	@echo "Create Namespace and install ArgoCD"
	@echo "-----------------------------------"
	kubectl create namespace "$(ARGO_NAMESPACE)" --dry-run=client -o yaml | kubectl apply -f -
	kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/$(ARGO_VERSION)/manifests/install.yaml

	@echo ""
	@echo "-------------------------------"
	@echo "Wait for Resources to get ready"
	@echo "-------------------------------"
	kubectl wait --for=condition=available deployment/argocd-dex-server -n "$(ARGO_NAMESPACE)" --timeout=120s
	kubectl wait --for=condition=available deployment/argocd-redis -n "$(ARGO_NAMESPACE)" --timeout=120s
	kubectl wait --for=condition=available deployment/argocd-repo-server -n "$(ARGO_NAMESPACE)" --timeout=120s
	kubectl wait --for=condition=available deployment/argocd-server  -n "$(ARGO_NAMESPACE)" --timeout=120s


	@echo ""
	@echo "#######################################################"
	@echo "ArgoCD Demo installed"
	@echo "- Get Admin Password: make argo-get-password"
	@echo "- Port-Forward ArgoCD: make port-forward-argocd"
	@echo "- Get Argo CLI: https://argo-cd.readthedocs.io/en/stable/cli_installation/"
	@echo "- Configure ArgoCD CLI (needs port-forward): make argo-configure-cli"
	@echo "- Install PodTatoHead via ArgoCD: make argo-install-podtatohead"	
	@echo "#######################################################"

.PHONY: argo-configure-cli
argo-configure-cli:
	@argocd login localhost:8080 --username admin --password $(ARGO_SECRET) --insecure

.PHONY: argo-get-password
argo-get-password:
	@echo $(ARGO_SECRET)

.PHONY: port-forward-argocd
port-forward-argocd:
	@echo ""
	@echo "Open ArgoCD in your Browser: http://localhost:8080"
	@echo "CTRL-c to stop port-forward"

	@echo ""
	kubectl port-forward svc/argocd-server -n "$(ARGO_NAMESPACE)" 8080:443

.PHONY: argo-install-podtatohead
argo-install-podtatohead:
	@echo ""
	kubectl apply -f config/app.yaml -n "$(ARGO_NAMESPACE)" 

.PHONY: uninstall
uninstall:
	kubectl delete -n $(ARGO_NAMESPACE) -f https://raw.githubusercontent.com/argoproj/argo-cd/$(ARGO_VERSION)/manifests/install.yaml --ignore-not-found=true
	kubectl delete namespace $(ARGO_NAMESPACE) --ignore-not-found=true
	@echo ""
	@echo "##########################"
	@echo "Argo Demo deleted"
	@echo "##########################"
