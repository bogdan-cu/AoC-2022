use std::{fs, io::Read};

fn main() {
    let mut file = match fs::File::open("/Users/bogdan/projects/AoC-2022/input/day1.txt") {
        Ok(file) => file,
        Err(err) => panic!("could not open file: {}", err),
    };
    let mut contents = String::new();
    match file.read_to_string(&mut contents) {
        Ok(_) => 1,
        Err(err) => panic!("could not read file as string: {}", err),
    };
    let max_calories = contents
        .trim()
        .split("\n\n")
        .map(|load| {
            let total_calories_per_person: usize = load
                .split('\n')
                .map(|item| match item.parse::<usize>() {
                    Ok(val) => val,
                    Err(err) => panic!("could not parse an item's caloric count of {}", err),
                })
                .sum();
            total_calories_per_person
        })
        .max()
        .unwrap();
    println!("the largest load has {} calories", max_calories);
}
