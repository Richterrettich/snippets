use std::borrow::Cow;


fn main() {
    let sources = vec![String::from("hallo"),String::from("welt")];
    let sources2 = vec!["bla","blubb"];

    test(&sources);
    test(&sources2);
}

fn cow_test<'a,T>(arg: T) where T: Into<Cow<'a,str>> {
    let raw: Cow<'a,str> = arg.into();
    println!("{}",raw);
}

fn test<T: AsRef<str>>(inp: &[T]) {
    for x in inp { println!("{} ", x.as_ref()) }

