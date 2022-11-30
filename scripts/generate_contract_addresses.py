import os
import json

client_common = os.environ.get('AVVY_CLIENT_COMMON', 'client-common')
contracts_path = os.path.join(client_common, 'contracts')
contract_files = os.listdir(contracts_path)

output = {
	'chains': {}
}
for ff in contract_files:
	chain_id = ff.replace('.json', '')
	with open(os.path.join(contracts_path, ff)) as f:
		data = json.loads(f.read())
		output['chains'][chain_id] = {
			'contracts': {
				contractValue: {
					'address': contractData['address']
				}
				for contractValue, contractData in data['contracts'].items()
			}
		}

output_path = os.path.join('avvy', 'contracts.json')
with open(output_path, 'w') as f:
	f.write(json.dumps(output))
