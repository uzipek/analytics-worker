import logging
import os
import json
import time
import hashlib

def load_config(file_path):
    try:
        with open(file_path, 'r') as f:
            return json.load(f)
    except FileNotFoundError:
        logging.error(f"Config file {file_path} not found")
        return None
    except json.JSONDecodeError:
        logging.error(f"Failed to parse config file {file_path}")
        return None

def hash_input(input_string):
    return hashlib.sha256(input_string.encode()).hexdigest()

def get_current_timestamp():
    return int(time.time())

def get_file_size(file_path):
    if not os.path.exists(file_path):
        return 0
    return os.path.getsize(file_path)

def create_dir(dir_path):
    if not os.path.exists(dir_path):
        os.makedirs(dir_path)