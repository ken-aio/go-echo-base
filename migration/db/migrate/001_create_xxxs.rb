class CreateXxxs < ActiveRecord::Migration
  def change
    create_table :xxxs do |t|
      t.string :name, null: false

      t.timestamps null: false
    end
  end
end
