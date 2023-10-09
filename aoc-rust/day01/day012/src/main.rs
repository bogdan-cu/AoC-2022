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
    let mut calories_per_person: Vec<usize> = contents
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
        .collect();
    calories_per_person.sort();
    calories_per_person.reverse();
    let top_three_loads = calories_per_person.into_iter().take(3).sum::<usize>();
    println!(
        "the strongest 3 lads are carrying a total of {} calories",
        top_three_loads
    );
}
