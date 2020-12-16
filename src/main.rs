use std::io;
use std::str;
use std::process::Command;
use tui::Terminal;
use tui::backend::TermionBackend;
//use tui::text::Text;
use termion::raw::IntoRawMode;
use tui::widgets::{List, ListItem, Block, Borders};
//use tui::widgets::Widget;
//use tui::layout::{Layout, Constraint, Direction};

fn main() -> Result<(), io::Error> {
    let stdout = io::stdout().into_raw_mode()?;
    let backend = TermionBackend::new(stdout);
    let mut terminal = Terminal::new(backend)?;
    terminal.hide_cursor()?;
    terminal.clear()?;

    //main loop 
    loop {
        terminal.draw(|f| {
            //create list of items in current directory
            let mut items = Vec::new();
            let ls_dir = Command::new("ls").output().expect("failed to read");
            let files = str::from_utf8(&ls_dir.stdout).unwrap();
            for (_, f) in files.split_whitespace().enumerate() {
                items.push(ListItem::new(f));
            }

            //draw list on terminal 
            let size = f.size();
            let list = List::new(items)
                .block(Block::default().title("Cowboy").borders(Borders::ALL));
            f.render_widget(list, size);
        })?
    }
}
