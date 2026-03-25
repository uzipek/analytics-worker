import logging
import os
import sys
from datetime import datetime
from typing import Optional

import pandas as pd
from dotenv import load_dotenv
from pydantic import BaseModel

logging.basicConfig(level=logging.INFO, format="%(asctime)s - %(levelname)s - %(message)s")
logger = logging.getLogger(__name__)

load_dotenv()

class Config(BaseModel):
    input_file: str = os.getenv("INPUT_FILE", "data/input.csv")
    output_dir: str = os.getenv("OUTPUT_DIR", "data/output")
    log_level: str = os.getenv("LOG_LEVEL", "INFO")

def validate_config(config: Config) -> bool:
    if not os.path.exists(config.input_file):
        logger.error(f"Input file not found: {config.input_file}")
        return False
    if not os.path.exists(config.output_dir):
        os.makedirs(config.output_dir)
    return True

def process_data(input_file: str, output_dir: str) -> Optional[str]:
    try:
        df = pd.read_csv(input_file)
        df["processed_at"] = datetime.now()
        output_file = os.path.join(output_dir, f"processed_{os.path.basename(input_file)}")
        df.to_csv(output_file, index=False)
        logger.info(f"Data processed and saved to: {output_file}")
        return output_file
    except Exception as e:
        logger.error(f"Error processing data: {e}")
        return None

def main():
    config = Config()
    if not validate_config(config):
        sys.exit(1)
    logger.setLevel(config.log_level)
    process_data(config.input_file, config.output_dir)

if __name__ == "__main__":
    main()