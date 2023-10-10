# generate mock files
mock:
	@paths="usecase"; \
	total=0; \
	for path in $$paths; do \
		total=$$((total + $$(find $$path -maxdepth 1 -mindepth 1 -type d | wc -l | tr -d '[:space:]'))); \
	done; \
	counter=0; \
	for path in $$paths; do \
		for dir in $$(find $$path -maxdepth 1 -mindepth 1 -type d); do \
			if [ -f $$dir/$$(basename $$dir).go ]; then \
				counter=$$((counter + 1)); \
				echo "Generating mock for $$dir ($$counter/$$total)"; \
				mockgen -source $$dir/$$(basename $$dir).go -destination $$dir/mock/$$(basename $$dir).go; \
			fi; \
		done; \
	done; \
	echo "Mock generation complete!"