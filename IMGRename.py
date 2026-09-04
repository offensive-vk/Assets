import os
import random

def get_valid_date():
    """Get and validate year and month input from user."""
    while True:
        try:
            year = int(input("Please enter a year (YYYY): "))
            if 1900 <= year <= 2100:
                break
            else:
                print("Please enter a valid year between 1900 and 2100.")
        except ValueError:
            print("Invalid format. Please enter a 4-digit year.")
    
    while True:
        try:
            month = int(input("Please enter a month (1-12): "))
            if 1 <= month <= 12:
                return year, month
            else:
                print("Please enter a valid month between 1 and 12.")
        except ValueError:
            print("Invalid format. Please enter a number between 1 and 12.")

def main():
    cwd = os.getcwd()
    print(f"Current directory: {os.path.abspath(cwd)}")
    
    # Find all .jpg image files
    jpg_files = [f for f in os.listdir(cwd) if f.lower().endswith('.jpg')]
    
    print(f"Found {len(jpg_files)} `.jpg` files.")
    
    if not jpg_files:
        print("No image files (.jpg) found. Exiting.")
        return
    
    # Get year and month
    year, month = get_valid_date()
    
    # Generate new names with unique random day and 6-digit numbers
    used_numbers = set()
    new_names = []
    for img in jpg_files:
        # Randomize day only
        day = random.randint(1, 31)
        
        # Generate unique 6-digit random number
        while True:
            rand_num = random.randint(100000, 999999)
            if rand_num not in used_numbers:
                used_numbers.add(rand_num)
                break
        new_name = f"IMG_{year}{month:02d}{day:02d}_{rand_num}.jpg"
        new_names.append((img, new_name))
    
    # Show preview of renames
    print("\nRenaming preview (first 5):")
    for i in range(min(5, len(jpg_files))):
        print(f"  {new_names[i][0]} → {new_names[i][1]}")
    
    # Confirm before renaming
    while True:
        response = input("\nProceed with renaming? (Y/n): ").lower()
        if response in ('y', ''):
            success_count = 0
            for original, new in new_names:
                try:
                    os.rename(os.path.join(cwd, original), os.path.join(cwd, new))
                    print(f"Renamed: {original} → {new}")
                    success_count += 1
                except Exception as e:
                    print(f"Error renaming {original}: {str(e)}")
            print("\nOperation complete!")
            print(f"Successfully renamed {success_count}/{len(jpg_files)} files.")
            return
        if response == 'n':
            print("Exiting...")
            return

if __name__ == "__main__":
    main()