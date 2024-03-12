import json

# Read data from JSON file
with open('pokemon.json', 'r') as file:
    pokemon_data = json.load(file)

with open('abilities.json', 'r') as file:
    ability_data = json.load(file)

with open('moves.json', 'r') as file:
    move_data = json.load(file)

with open('items.json', 'r') as file:
    item_data = json.load(file)

def generate_insert_statements(pokemon_data, sql_file_path):
    with open(sql_file_path, 'w') as sql_file:
        for pokemon in pokemon_data:
            english_name = pokemon['name']['english']
            type1 = pokemon['type'][0]
            type2 = pokemon['type'][1] if len(pokemon['type']) > 1 else ''
            stats = json.dumps(pokemon['base'])
            
            # Escape single quotes in the English name
            # Looking at you Farfetch'd and Sirfetch'd
            english_name_escaped = english_name.replace("'", "''")
            
            # Write SQL insert statement to the file
            sql_file.write(f"INSERT INTO pokemon.pokemon (name, type1, type2, stats) VALUES ('{english_name_escaped}', '{type1}', '{type2}', '{stats}');\n")

def generate_ability_insert_statements(ability_data, sql_file_path):
    with open(sql_file_path, 'w') as sql_file:
        for ab in ability_data:
            ability_name = ab['ability_name']
            ability_description = ab['ability_description']
            
            desc_escaped = ability_description.replace("'", "''")

            sql_file.write(f"INSERT INTO pokemon.abilities (ability_name, ability_description) VALUES ('{ability_name}', '{desc_escaped}');\n")

def generate_move_insert_statements(move_data, sql_file_path):
    with open(sql_file_path, 'w') as sql_file:
        for move in move_data:
            move_name = move['move_name']
            move_type = move['move_type']
            move_power = move['move_power'] if 'move_power' in move else 0
            move_category = move['move_category']
            move_accuracy = move['move_accuracy']
            move_pp = move['move_pp']
            move_effect = move['effect'] if 'effect' in move else 'None'
            move_effect_chance = move['effect']['chance'] if 'move_effect_chance' in move else 0
            move_status_condition = move['effect']['statusCondition'] if 'move_status_condition' in move else ''
            move_priority = move['move_priority'] if 'move_priority' in move else 0
            
            sql_file.write(f"INSERT INTO moves.moves (move_name, move_type, move_power, move_category, move_accuracy, move_pp, move_effect, move_effect_chance, move_status_condition, move_priority) VALUES ('{move_name}', '{move_type}', {move_power}, '{move_category}', {move_accuracy}, {move_pp}, '{move_effect}', {move_effect_chance}, '{move_status_condition}', {move_priority});\n")

def generate_item_insert_statements(item_data, sql_file_path):
    with open(sql_file_path, 'w') as sql_file:
        for item in item_data:
            item_name = item['name']['english']
            item_description = item['description']
            item_price = item['price']

            desc_escaped = item_description.replace("'", "''")

            sql_file.write(f"INSERT INTO items.items (item_name, item_description, item_price) VALUES ('{item_name}', '{desc_escaped}', {item_price});\n")

def main():
    sql_file_path = 'pokemon_insert.sql'
    generate_insert_statements(pokemon_data, sql_file_path)
    print(f"Generated SQL insert statements for Pokemon data at {sql_file_path}")
    generate_ability_insert_statements(ability_data, 'abilities_insert.sql')
    print(f"Generated SQL insert statements for Ability data at abilities_insert.sql")
    generate_move_insert_statements(move_data, 'moves_insert.sql')
    print(f"Generated SQL insert statements for Move data at moves_insert.sql")
    generate_item_insert_statements(item_data, 'items_insert.sql')
    print(f"Generated SQL insert statements for Item data at items_insert.sql")

main()