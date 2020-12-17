use tui::widgets::{List, ListItem, Block, Borders};

//Parent represents the parent direcotry, current is the current direcotry, and Child is the
//contents of whatever the cursor is highlighting
pub enum PaneType {
    Parent,
    Current,
    Child,
}

//Pane is used to represent a singular block in the terminal 
pub struct Pane<'a> {
    items: Vec<ListItem<'a>>,
    pub ptype: PaneType,
}

impl Pane<'_> {
    //Create a new pane given a &str, which is then split into a list 
    pub fn new(content: &str, ptype: PaneType) -> Pane {
        let mut items = Vec::new();
        for (_, f) in content.split_whitespace().enumerate() {
            items.push(ListItem::new(f));
        }
        Pane {
            items,
            ptype,
        }

    } 

    //Render the pane
    pub fn visualize(
        self,
        t: &mut tui::Frame<tui::backend::TermionBackend<termion::raw::RawTerminal<std::io::Stdout>>>,
        rect: tui::layout::Rect,
    ){
        let list = List::new(self.items)
            .block(Block::default().title("Cowboy").borders(Borders::ALL));
        t.render_widget(list, rect);

    }
}
