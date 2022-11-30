import json
import os

with open('./client-common/contracts/43114.json') as f:
	data = json.loads(f.read())

for key, contract in data['contracts'].items():
	with open(f'./abi/{key}.json', 'w') as f:
		f.write(json.dumps(contract['abi']))
